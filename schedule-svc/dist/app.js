"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const cors_1 = __importDefault(require("cors"));
const http_1 = __importDefault(require("http"));
const socket_1 = require("./src/socket");
require("./src/config/firebase-config");
require("./src/config/google-config");
require("dotenv").config();
const { errorHandler } = require("./src/handler/errorHandler");
const app = (0, express_1.default)();
app.use(express_1.default.json());
app.use(express_1.default.urlencoded({ extended: true }));
app.use((0, cors_1.default)());
// Import route
const scheduleGARoute_1 = __importDefault(require("./src/routes/scheduleGARoute"));
const pointDefRoute_1 = __importDefault(require("./src/routes/pointDefRoute"));
const notificationRoute_1 = __importDefault(require("./src/routes/notificationRoute"));
// Mouting the route
app.use("/api/schedule-report/", scheduleGARoute_1.default);
app.use("/api/point-def", pointDefRoute_1.default);
app.use("/api/notification", notificationRoute_1.default);
const server = http_1.default.createServer(app);
(0, socket_1.configureSocket)(server);
const PORT = process.env.PORT || 5000;
app.all("*", (req, res, next) => {
    const err = new Error("The route can not be found");
    next(err);
});
app.use(errorHandler);
server.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
const path_1 = __importDefault(require("path"));
const grpc = __importStar(require("@grpc/grpc-js"));
const protoLoader = __importStar(require("@grpc/proto-loader"));
const GAController_1 = require("./src/controllers/GAController");
const notificationController_1 = require("./src/controllers/notificationController");
const pointDefController_1 = require("./src/controllers/pointDefController");
const PORT_PROTO = 9091;
const PROTO_FILE = "./src/proto/schedule.proto";
const packageDef = protoLoader.loadSync(path_1.default.resolve(__dirname, PROTO_FILE));
const grpcObj = grpc.loadPackageDefinition(packageDef);
const schedulePkg = grpcObj.schedule;
const serverGrpc = getServer();
serverGrpc.bindAsync(`0.0.0.0:${PORT_PROTO}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
        console.error(err);
        return;
    }
    console.log(`Server gRPC is running on ${port}`);
    serverGrpc.start();
});
function getServer() {
    const server = new grpc.Server();
    server.addService(schedulePkg.v1.ScheduleService.service, {
        GetSchedules: (req, res) => {
            const schedule = GAController_1.GAScheduleController.getSchedule(req);
            schedule.then((response) => {
                res(null, {
                    id: response === null || response === void 0 ? void 0 : response.id,
                    thesis: response === null || response === void 0 ? void 0 : response.thesis,
                });
            });
        },
        CreateSchedule: (req, callback) => {
            GAController_1.GAScheduleController.scheduleGAThesisDefense(req);
            let schedule = GAController_1.GAScheduleController.getSchedule(req);
            schedule.then((response) => {
                callback(null, {
                    id: response === null || response === void 0 ? void 0 : response.id,
                    thesis: response === null || response === void 0 ? void 0 : response.thesis,
                });
            });
        },
        CreateNotification: (req, callback) => {
            let notis = notificationController_1.notificationController.createNotification(req);
            notis.then((response) => {
                callback(null, {
                    notification: response === null || response === void 0 ? void 0 : response.notification,
                    message: response === null || response === void 0 ? void 0 : response.message,
                    notifications: response === null || response === void 0 ? void 0 : response.notifications,
                });
            });
        },
        CreateOrUpdatePointDef: (req, callback) => {
            let res = pointDefController_1.PointDefController.createOrUpdatePointDef(req);
            res.then((point) => {
                callback(null, {
                    point: point,
                    message: "PointDef has been created",
                });
            });
        },
        GetAllPointDefs: (req, callback) => {
            let res = pointDefController_1.PointDefController.getAllPointDef(req);
            res.then((point) => {
                console.log(point);
                callback(null, {
                    point: point,
                });
            });
        },
        UpdatePointDef: (req, callback) => {
            let res = pointDefController_1.PointDefController.updatePointDef(req);
            res.then((point) => {
                console.log(point);
                callback(null, {
                    point: point
                });
            });
        },
        DeletePointDef: (req, callback) => {
            let res = pointDefController_1.PointDefController.deletePointDef(req);
            res.then((isDelete) => {
                if (isDelete) {
                    callback(null, {
                        message: "point has been deleted"
                    });
                }
                else {
                    callback(null, {
                        message: "some error occurred"
                    });
                }
            });
        }
    });
    return server;
}
