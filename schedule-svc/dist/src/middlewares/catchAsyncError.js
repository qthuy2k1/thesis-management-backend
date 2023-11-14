"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
module.exports =
    (theFunc) => (req, res, next) => {
        Promise.resolve(theFunc(req, res, next)).catch(next);
    };
