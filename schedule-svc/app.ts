import express, { response } from "express";
import cors from "cors";
import http from "http";
import { Request, Response, NextFunction } from "express";
import { configureSocket } from "./src/socket";
require("./src/config/firebase-config");
require("./src/config/google-config");
require("dotenv").config();
const { errorHandler } = require("./src/handler/errorHandler");
const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

// Import route
import scheduleGARoute from "./src/routes/scheduleGARoute";
import pointDefRoute from "./src/routes/pointDefRoute";
import notificationRoute from "./src/routes/notificationRoute";

// Mouting the route
app.use("/api/schedule-report/", scheduleGARoute);
app.use("/api/point-def", pointDefRoute);
app.use("/api/notification", notificationRoute);

const server = http.createServer(app);
configureSocket(server);
const PORT = process.env.PORT || 5000;
app.all("*", (req: Request, res: Response, next: NextFunction) => {
  const err = new Error("The route can not be found");
  next(err);
});
app.use(errorHandler);
server.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});

import path from "path";
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { ProtoGrpcType } from "./src/proto/schedule";
import { GAScheduleController } from "./src/controllers/GAController";
import { notificationController } from "./src/controllers/notificationController";
import { PointDefController } from "./src/controllers/pointDefController";
import { ExerciseController } from "./src/controllers/exerciseController";

const PORT_PROTO = 9091;
const PROTO_FILE = "./src/proto/schedule.proto";

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE));
const grpcObj = grpc.loadPackageDefinition(
  packageDef
) as unknown as ProtoGrpcType;

const schedulePkg = grpcObj.schedule;

const serverGrpc = getServer();
serverGrpc.bindAsync(
  `0.0.0.0:${PORT_PROTO}`,
  grpc.ServerCredentials.createInsecure(),
  (err, port) => {
    if (err) {
      console.error(err);
      return;
    }
    console.log(`Server gRPC is running on ${port}`);
    serverGrpc.start();
  }
);

function getServer() {
  const server = new grpc.Server();
  server.addService(schedulePkg.v1.ScheduleService.service, {
    GetSchedules: (req: any, res: any) => {
      const schedule = GAScheduleController.getSchedule(req);
      schedule.then((response) => {
        res(null, {
          id: response?.id,
          thesis: response?.thesis,
        });
      });
    },
    CreateSchedule: (req: any, callback: any) => {
      GAScheduleController.scheduleGAThesisDefense(req);
      let schedule = GAScheduleController.getSchedule(req);
      schedule.then((response) => {
        callback(null, {
          id: response?.id,
          thesis: response?.thesis,
        });
      });
    },
    CreateNotification: (req: any, callback: any) => {
      let notis = notificationController.createNotification(req);

      notis.then((response) => {
        callback(null, {
          notification: response?.notification,
          message: response?.message,
          notifications: response?.notifications,
        });
      });
    },
    CreateOrUpdatePointDef: (req: any, callback: any) => {
      let point = PointDefController.createOrUpdatePointDef(req);

      point.then((response) => {
        callback(null, {
          point: response,
          message: "PointDef has been created",
        });
      });
    },
    GetAllPointDefs: (req: any, callback: any) => {
      let points = PointDefController.getAllPointDef(req);

      points.then((response) => {
        console.log(response)
        callback(null, {
          points: response,
        });
      });
    },
  });

  return server;
}
