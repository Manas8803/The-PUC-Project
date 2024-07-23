const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
  const connectionId = event.requestContext.connectionId;

  try {
     const params = {
    TableName: process.env.RTO_OFFICE_TABLE_ARN,
    FilterExpression: "contains(connection_id, :connectionValue)",
    ExpressionAttributeValues: {
        ":connectionValue": connectionId // The value you are looking for in the connection_id list
    }};
    // First, find the item with the given connectionId
    const data = await ddb
      .scan(params)
      .promise();

    if (data.Items.length === 0) {
      return {
        statusCode: 404,
        body: JSON.stringify({
          message: "No item found with the given connectionId",
        }),
      };
    }

    const itemToUpdate = data.Items[0];
    
    const updatedConnectionIds = itemToUpdate.connection_id.filter(
      (id) => id !== connectionId
    );

    // Update the item with the new connectionIds array
    await ddb
      .put({
        TableName: process.env.RTO_OFFICE_TABLE_ARN,
        Item: {
          office_name: itemToUpdate.office_name,
          connection_id: updatedConnectionIds,
        },
      })
      .promise();

    return {
      statusCode: 200,
      body: JSON.stringify({ message: "Disconnected successfully" }),
    };
  } catch (err) {
    console.error("Error disconnecting:", err);
    return {
      statusCode: 500,
      body: JSON.stringify({ error: "Error disconnecting" }),
    };
  }
};