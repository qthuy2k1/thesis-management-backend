"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.errorHandler = void 0;
const errorHandler = (error, req, res, next) => {
    const errorMessage = error.message || "Unknown error";
    const statusCode = error.statusCode || 400;
    res.status(400).json({
        status: "Fail to call API",
        message: errorMessage,
    });
};
exports.errorHandler = errorHandler;
