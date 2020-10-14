const path = require('path');
const express = require('express');
const bodyParser = require('body-parser');
const notifier = require('node-notifier');

const app = express();
const port = process.env.PORT || 3100;

app.use(bodyParser.json());

app.get('/health', (req, res) => res.status(200).send('ok'));
app.post('/notify', (req, res) => {
    notify(req.body, reply => res.send(reply));
});

app.listen(port, () => console.log(`server running on port ${port}`));


const notify = ({ title, message }, callback) => {
    notifier.notify(
        {
            title: title || "Unknown Title",
            message: message || "Unknown message",
            sound: true,
            wait: true,
            reply: true,
            icon: path.join(__dirname, "reminder.png"),
            closeLabel: "Completed",
            timeout: 15
        },
        (err, response, reply) => {
            callback(reply)
        }
    )
};