"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PointDefController = void 0;
const pointDefModel_1 = require("../models/pointDefModel");
exports.PointDefController = {
    async createOrUpdatePointDef(req) {
        const point = req.request.point;
        console.log(point.assesses);
        // let input = {
        // }
        try {
            return await pointDefModel_1.PointDefModel.createOrUpdatePointDef(point);
        }
        catch (err) {
            console.log(err);
        }
        return;
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
            const id = req.request.id;
            const points = await pointDefModel_1.PointDefModel.getAllPointDef(id);
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
            return;
        }
    },
    async updatePointDef(req) {
        const point = req.request.point;
        const id = req.request.id;
        try {
            await pointDefModel_1.PointDefModel.updatePointDef({ id, ...point });
            // res.status(200).json({ id, ...point });
            return { id, ...point };
        }
        catch (err) {
            console.error(err);
            // res.status(500);
        }
    },
    async deletePointDef(req) {
        const id = req.request.id;
        try {
            const points = await pointDefModel_1.PointDefModel.deletePointDef(id);
            if (!points) {
                // res.status(404);
                return false;
            }
            // return res.status(200).json(points);
            return true;
        }
        catch (err) {
            console.error(err);
            // res.status(500);
            return false;
        }
    },
};
