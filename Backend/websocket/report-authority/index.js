const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
  let connections;
  const body = JSON.parse(event.body);
  const rtoOfficeName = body.data.office_name;

  if (!rtoOfficeName) {
    return {
      statusCode: 400,
      message: "Missing office_name parameter in the query string.",
    };
  }

  try {
    const getParams = {
      TableName: process.env.RTO_OFFICE_TABLE_ARN,
      Key: {
        office_name: rtoOfficeName,
      },
    };

    const getResult = await ddb.get(getParams).promise();
    connections = getResult.Item ? getResult.Item.connection_id : null;
  } catch (err) {
    return {
      statusCode: 500,
      message: "Internal Server Error: Error in fetching from DB.",
    };
  }

  if (!connections || connections.length === 0) {
    return {
      statusCode: 404,
      message: `No connections found for office_name: ${rtoOfficeName}`,
    };
  }

  const callbackAPI = new AWS.ApiGatewayManagementApi({
    apiVersion: "2018-11-29",
    endpoint: event.requestContext.domainName + "/" + event.requestContext.stage,
  });

  let errors = [];
  
  for (let conn of connections) {
    const connectionId = conn;
    if (connectionId !== event.requestContext.connectionId) {
      try {
        await callbackAPI
          .postToConnection({ ConnectionId: connectionId, Data: JSON.stringify(body.data) })
          .promise();
      } catch (e) {
        console.log(`Error posting to connection ${connectionId}:`, e);
        errors.push(connectionId);
      }
    }
  }

  if (errors.length > 0) {
    return {
      statusCode: 500,
      message: `Failed to post to some connections: ${errors.join(", ")}`,
    };
  }

  return { statusCode: 200 };
};
