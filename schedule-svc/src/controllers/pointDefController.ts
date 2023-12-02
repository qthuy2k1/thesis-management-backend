import { Request, Response } from "express";
import { PointDefModel } from "../models/pointDefModel";
import { IPointDefObject } from "../interface/pointDef";

export const PointDefController = {
  async createOrUpdatePointDef(req: any): Promise<IPointDefObject | undefined>{
    const point = req.request.point;

    console.log(point.assesses)
    
    // let input = {

    // }
    try {
      return await PointDefModel.createOrUpdatePointDef(point);
    } catch (err) {
      console.log(err);
    }
    return
  },

  async getPointDef(req: Request, res: Response) {
    const id = req.params.id;
    try {
      const point = await PointDefModel.getPointDef(id);
      if (!point) {
        res.status(404);
        return;
      }
      res.status(200).json(point);
    } catch (err) {
      console.error(err);
      res.status(500);
    }
  },

  async getPointDefForLecturer(req: Request, res: Response) {
    const { studefId } = req.params;
    const { lecId } = req.params;
    try {
      const point = await PointDefModel.getPointDefForLecturer(studefId, lecId);
      if (!point) {
        res.status(404);
        return;
      }
      res.status(200).json(point);
    } catch (err) {
      console.error(err);
      res.status(500);
    }
  },

  async getAllPointDef(req: any): Promise<IPointDefObject[] | undefined> {
    try {
      const id = req.request.id;
      const points = await PointDefModel.getAllPointDef(id);
      if (!points) {
        // res.status(404).json("PointDef is empty");
        return;
      }
      // res.status(200).json(points);

      return points;
    } catch (err) {
      console.error(err);
      // res.status(500);
      return
    }
  },

  // async updatePointDef(req: Request): Promise<IPointDefObject[] | undef> {
  //   const point = req.body;
  //   const id = req.params.id;
  //   try {
  //     await PointDefModel.updatePointDef({ id, ...point });
  //     res.status(200).json({ id, ...point });
  //   } catch (err) {
  //     console.error(err);
  //     res.status(500);
  //   }
  // },

  async deletePointDef(req: Request, res: Response) {
    const id = req.params.id;
    try {
      const points = await PointDefModel.deletePointDef(id);
      if (!points) {
        res.status(404);
        return false;
      }
      return res.status(200).json(points);
    } catch (err) {
      console.error(err);
      res.status(500);
    }
  },
};
