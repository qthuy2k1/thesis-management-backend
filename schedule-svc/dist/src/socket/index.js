"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.configureSocket = void 0;
const socket_io_1 = require("socket.io");
const axios_1 = __importDefault(require("axios"));
function configureSocket(server) {
    const io = new socket_io_1.Server(server, {
        cors: {
            origin: "http://localhost:3000",
            methods: ["GET", "POST"],
        },
    });
    // EVENT SOCKET.IO
    let onlineUsers = [];
    const addNewUser = (user) => {
        const foundIndex = onlineUsers.findIndex((prevUser) => prevUser.id === user.id);
        if (foundIndex !== -1) {
            onlineUsers.splice(foundIndex, 1);
        }
        if (!onlineUsers.some((prevUser) => prevUser.socketId === user.socketId)) {
            onlineUsers.push(user);
        }
    };
    const removeUser = (socketId) => {
        return onlineUsers.filter((nextUser) => nextUser.socketId !== socketId);
    };
    const getUser = (id) => {
        return onlineUsers.find((foundUser) => foundUser.id === id);
    };
    io.on("connect", (socket) => {
        socket.on("newUser", (user) => {
            addNewUser(user);
            if (onlineUsers.some((prevUser) => prevUser.id === user.id)) {
                axios_1.default
                    .get(`http://localhost:5000/api/notification/${user.id}`)
                    .then((response) => {
                    io.to(user.socketId || "").emit("getAllNotifications", response.data);
                })
                    .catch((error) => {
                    console.error("Error", error);
                });
            }
        });
        socket.on("sendNotification", ({ senderUser, receiverAuthor, type }) => {
            const receiver = getUser(receiverAuthor.id || "");
            if (receiver && (receiver === null || receiver === void 0 ? void 0 : receiver.socketId)) {
                io.to(receiver === null || receiver === void 0 ? void 0 : receiver.socketId).emit("updateNotifications", {
                    senderUser,
                    receiverAuthor,
                    type,
                });
                axios_1.default
                    .post("http://localhost:5000/api/notification", {
                    senderUser,
                    receiverAuthor,
                    type,
                }, {
                    headers: {
                        "Content-Type": "application/json",
                    },
                })
                    .then((response) => {
                    io.to(receiver.socketId || "").emit("updateNotification", response.data.notifications);
                })
                    .catch((error) => {
                    console.error("Error", error);
                });
            }
        });
        socket.on("disconnect", () => {
            removeUser(socket.id);
        });
    });
    return io;
}
exports.configureSocket = configureSocket;
