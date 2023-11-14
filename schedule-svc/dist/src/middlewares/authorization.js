"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const admin = require("../config/firebase-config");
class CheckAuthorization {
    static async decodeToken(req, res, next) {
        var _a;
        const token = (_a = req.headers.authorization) === null || _a === void 0 ? void 0 : _a.split(" ")[1];
        try {
            const decodeValue = await admin.auth().verifyIdToken(token);
            if (decodeValue) {
                return next();
            }
            return res.json({ message: "Unauthorization" });
        }
        catch (error) {
            return res.json({ message: "Internal Error" });
        }
    }
}
exports.default = CheckAuthorization.decodeToken;
