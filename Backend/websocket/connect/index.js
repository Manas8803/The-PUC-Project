const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
	const rtoOfficeName = event.queryStringParameters
		? event.queryStringParameters.office_name
		: null;
	if (!rtoOfficeName) {
		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Connected successfully" }),
		};
	}

	//* Checks whether the office_name exists in the user table
	try {
		const params = {
			TableName: process.env.USER_TABLE_ARN,
			FilterExpression: "office_name = :officeName",
			ExpressionAttributeValues: {
				":officeName": rtoOfficeName,
			},
		};

		const data = await ddb.scan(params).promise();

		if (data.Items.length === 0) {
			return {
				statusCode: 404,
				body: JSON.stringify({
					message: `Office name '${rtoOfficeName}' is not registered`,
				}),
			};
		}
	} catch (err) {
		console.error("Error checking office name:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error checking office name" }),
		};
	}

	//* Add an entry to the table with single connection across office_name
	try {

    // First, retrieve the existing connectionIds (if any) for the given office_name
    const existingData = await ddb
      .get({
        TableName: process.env.RTO_OFFICE_TABLE_ARN,
        Key: { office_name: rtoOfficeName },
      })
      .promise();

    let connectionIds = existingData.Item?.connection_id || [];
    if (connectionIds.length >= 3) {
      return {
        statusCode: 400,
        body: JSON.stringify({
          error: "Maximum number of connections reached",
        }),
      };
    }

    // Add the new connectionId to the array
    connectionIds.push(event.requestContext.connectionId);

    // Update or create the item with the new connectionIds array
    await ddb
      .put({
        TableName: process.env.RTO_OFFICE_TABLE_ARN,
        Item: {
          office_name: rtoOfficeName,
          connection_id: connectionIds,
        },
      })
      .promise();

    return {
      statusCode: 200,
      body: JSON.stringify({ message: "Connected successfully" }),
    };
  } catch (err) {
    console.error("Error connecting:", err);
    return {
      statusCode: 500,
      body: JSON.stringify({ error: "Error connecting" }),
    };
  }
};