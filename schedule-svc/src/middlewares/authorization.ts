import { Request, Response, NextFunction } from "express";
const admin = require("../config/firebase-config");

class CheckAuthorization {
  static async decodeToken(req: Request, res: Response, next: NextFunction) {
    const token = req.headers.authorization?.split(" ")[1];
    try {
      const decodeValue = await admin.auth().verifyIdToken(token);
      if (decodeValue) {
        return next();
      }
      return res.json({ message: "Unauthorization" });
    } catch (error) {
      return res.json({ message: "Internal Error" });
    }
  }
}

export default CheckAuthorization.decodeToken;
