"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TopicController = void 0;
const topicModel_1 = require("../models/topicModel");
exports.TopicController = {
    async createTopic(req, res) {
        const topic = req.body;
        try {
            await topicModel_1.TopicModel.createTopic(topic);
            res.status(200).json({ topic, message: "Topic has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getTopic(req, res) {
        const id = req.params.id;
        try {
            const topic = await topicModel_1.TopicModel.getTopic(id);
            if (!topic) {
                res.status(404);
                return;
            }
            res.status(200).json(topic);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllTopic(req, res) {
        try {
            const topics = await topicModel_1.TopicModel.getAllTopic();
            if (!topics) {
                res.status(404).json("Topic is empty");
                return;
            }
            res.status(200).json(topics);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateTopic(req, res) {
        const topic = req.body;
        const id = req.params.id;
        try {
            await topicModel_1.TopicModel.updateTopic({ id, ...topic });
            res.status(200).json({ id, ...topic });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteTopic(req, res) {
        const id = req.params.id;
        try {
            const topics = await topicModel_1.TopicModel.deleteTopic(id);
            if (!topics) {
                res.status(404);
                return false;
            }
            return res.status(200).json(topics);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
