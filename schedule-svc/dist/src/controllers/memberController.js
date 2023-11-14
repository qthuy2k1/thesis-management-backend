"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MemberController = void 0;
const memberModel_1 = require("../models/memberModel");
exports.MemberController = {
    async createMember(req, res) {
        const member = req.body;
        try {
            await memberModel_1.MemberModel.createMember(member);
            res.status(200).json({ member, message: "Member has been created" });
        }
        catch (err) {
            res.status(400).json({ message: err });
        }
    },
    async getMember(req, res) {
        const id = req.params.id;
        try {
            const member = await memberModel_1.MemberModel.getMember(id);
            if (!member) {
                res.status(404);
                return;
            }
            res.status(200).json(member);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllMember(req, res) {
        try {
            const members = await memberModel_1.MemberModel.getAllMember();
            if (!members) {
                res.status(404).json("Member is empty");
                return;
            }
            res.status(200).json(members);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllMemberClassroom(req, res) {
        try {
            const id = req.params.id;
            const members = await memberModel_1.MemberModel.getAllMemberClassroom(id);
            if (!members) {
                res.status(404).json("Member is empty");
                return;
            }
            res.status(200).json(members);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateMember(req, res) {
        const member = req.body;
        const id = req.params.id;
        try {
            await memberModel_1.MemberModel.updateMember({ id, ...member });
            res.status(200).json({ id, ...member });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteMember(req, res) {
        const id = req.params.id;
        try {
            const members = await memberModel_1.MemberModel.deleteMember(id);
            if (!members) {
                res.status(404);
                return false;
            }
            return res.status(200).json(members);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
