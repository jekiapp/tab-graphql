const express = require('express');
const graphqlHTTP = require('express-graphql');

const app = express();
app.use(express.static(__dirname));

app.listen(3000, function() {
  const port = this.address().port;
  console.log(`Started on http://localhost:${port}/`);
});
