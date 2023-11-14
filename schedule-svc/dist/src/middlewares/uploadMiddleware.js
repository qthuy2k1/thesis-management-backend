"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const multer_1 = __importDefault(require("multer"));
const fs_1 = __importDefault(require("fs"));
const storage = multer_1.default.diskStorage({
    destination: function (req, file, cb) {
        const post = req.body;
        const uploadPath = `src/uploads/${post.uid}`;
        fs_1.default.access(uploadPath, (error) => {
            if (error) {
                fs_1.default.mkdir(uploadPath, { recursive: true }, (error) => {
                    if (error) {
                        console.error("Error creating upload directory:", error);
                        cb(error, "err");
                    }
                    else {
                        cb(null, uploadPath);
                    }
                });
            }
            else {
                cb(null, uploadPath);
            }
        });
    },
    filename: function (req, file, cb) {
        cb(null, file.originalname);
    },
});
const uploadMiddleware = (0, multer_1.default)({ storage: storage });
exports.default = uploadMiddleware;
