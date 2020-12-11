const express = require("express");

const app = express();

app.get("/", (req, res) => {
  res.json({ status: "running", language: "JS" });
});

app.get("/hello", (req, res) => {
    res.json({...req.query})
})

app.listen(4000, () => {
  console.log("app running on port 4000");
});
