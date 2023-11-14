"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PointDefController = void 0;
const pointDefModel_1 = require("../models/pointDefModel");
exports.PointDefController = {
    async createOrUpdatePointDef(req) {
        const point = req.request.point;
        try {
            return await pointDefModel_1.PointDefModel.createOrUpdatePointDef(point);
            // res.status(200).json({ point, message: "PointDef has been created" });
        }
        catch (err) {
            // res.status(400).json({ message: err });
            console.log(err);
        }
    },
    async getPointDef(req, res) {
        const id = req.params.id;
        try {
            const point = await pointDefModel_1.PointDefModel.getPointDef(id);
            if (!point) {
                res.status(404);
                return;
            }
            res.status(200).json(point);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getPointDefForLecturer(req, res) {
        const { studefId } = req.params;
        const { lecId } = req.params;
        try {
            const point = await pointDefModel_1.PointDefModel.getPointDefForLecturer(studefId, lecId);
            if (!point) {
                res.status(404);
                return;
            }
            res.status(200).json(point);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllPointDef(req) {
        try {
            const points = await pointDefModel_1.PointDefModel.getAllPointDef();
            if (!points) {
                // res.status(404).json("PointDef is empty");
                return;
            }
            // res.status(200).json(points);
            return points;
        }
        catch (err) {
            console.error(err);
            // res.status(500);
        }
    },
    async updatePointDef(req, res) {
        const point = req.body;
        const id = req.params.id;
        try {
            await pointDefModel_1.PointDefModel.updatePointDef({ id, ...point });
            res.status(200).json({ id, ...point });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deletePointDef(req, res) {
        const id = req.params.id;
        try {
            const points = await pointDefModel_1.PointDefModel.deletePointDef(id);
            if (!points) {
                res.status(404);
                return false;
            }
            return res.status(200).json(points);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
