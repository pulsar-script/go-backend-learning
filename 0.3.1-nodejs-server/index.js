const express = require('express')
const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
    res.status(200).send("welcome to Node js dummy server made for Go lang");
});

app.get("/hello", (req, res) => {
    res.status(200).json({ message: "Hello message from Node js server" });
});


app.post("/send/message", (req, res) => {

    let myJson = req.body;
    console.log(myJson);
    res.status(200).send(myJson);

});

app.post("/send/form", (req, res) => {

    console.log("req.body => ", req.body);
    jsonData = JSON.stringify(req.body)
    console.log(" jsonData => ", jsonData);
    res.status(200).send(jsonData);

});

app.listen(port, () => {
    console.log("server is listing at port 3000");
});