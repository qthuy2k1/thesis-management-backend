"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuthController = void 0;
const authModel_1 = require("../models/authModel");
// const hashPassword = async (password: string): Promise<string> => {
//   const saltRounds = 10;
//   const hashedPassword = await bcrypt.hash(password, saltRounds);
//   const truncatedHash = hashedPassword.slice(0, 10);
//   return truncatedHash;
// };
exports.AuthController = {
    async createAuth(req, res) {
        const { password, ...auth } = req.body;
        // const hashPasswordAuth = await hashPassword(password);
        try {
            await authModel_1.AuthModel.createAuth({ password: password, ...auth });
            res.status(200).json({ auth, message: "Auth has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getAuth(req, res) {
        const id = req.params.id;
        try {
            const auth = await authModel_1.AuthModel.getAuth(id);
            if (!auth) {
                res.status(404);
                return;
            }
            res.status(200).json(auth);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllAuth(req, res) {
        try {
            const auths = await authModel_1.AuthModel.getAllAuth();
            if (!auths) {
                res.status(404).json("Auth is empty");
                return;
            }
            res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllLecturer(req, res) {
        try {
            const auths = await authModel_1.AuthModel.getAllLecturer();
            if (!auths) {
                res.status(404).json("Auth is empty");
                return;
            }
            res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateAuth(req, res) {
        const auth = req.body;
        const id = req.params.id;
        try {
            await authModel_1.AuthModel.updateAuth({ id, ...auth });
            res.status(200).json({ id, ...auth });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteAuth(req, res) {
        const id = req.params.id;
        try {
            const auths = await authModel_1.AuthModel.deleteAuth(id);
            if (!auths) {
                res.status(404);
                return false;
            }
            return res.status(200).json(auths);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async checkStatusSubscribe(req, res) {
        const id = req.params.id;
        try {
            const auths = await authModel_1.AuthModel.checkStatusSubscribe(id);
            if (Array.isArray(auths)) {
                return res.status(200).json(auths);
            }
            else {
                return res.status(200).json({ status: "NO_SUBSCRIBE" });
            }
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async checkAuthRoleForClassroomState(req, res) {
        const auth = req.body;
        try {
            const auths = await authModel_1.AuthModel.checkAuthRoleForClassroomState(auth);
            if (auths) {
                return res.status(200).json(auths);
            }
            else {
                return res.status(404).json("Classroom is not found");
            }
        }
        catch (err) {
            console.error(err);
            return res.status(500);
        }
    },
    async unsubscribeState(req, res) {
        const id = req.params.id;
        try {
            const auths = await authModel_1.AuthModel.unsubscribeState(id);
            return res
                .status(200)
                .json({ auth: auths, message: "Unsubscribe successfully" });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
