import { Server, Socket } from "socket.io";
import http from "http";
import axios from "axios";
import { IAuthObject } from "../interface/auth";
import { INotificationObject } from "../interface/notification";

export function configureSocket(server: http.Server): Server {
  const io = new Server(server, {
    cors: {
      origin: "http://localhost:3000",
      methods: ["GET", "POST"],
    },
  });
  // EVENT SOCKET.IO

  let onlineUsers: IAuthObject[] = [];

  const addNewUser = (user: IAuthObject) => {
    const foundIndex = onlineUsers.findIndex(
      (prevUser: IAuthObject) => prevUser.id === user.id
    );

    if (foundIndex !== -1) {
      onlineUsers.splice(foundIndex, 1);
    }

    if (
      !onlineUsers.some(
        (prevUser: IAuthObject) => prevUser.socketId === user.socketId
      )
    ) {
      onlineUsers.push(user);
    }
  };
  const removeUser = (socketId: string) => {
    return onlineUsers.filter((nextUser) => nextUser.socketId !== socketId);
  };

  const getUser = (id: string) => {
    return onlineUsers.find((foundUser) => foundUser.id === id);
  };
  io.on("connect", (socket: Socket) => {
    socket.on("newUser", (user: IAuthObject) => {
      addNewUser(user);
      if (
        onlineUsers.some((prevUser: IAuthObject) => prevUser.id === user.id)
      ) {
        axios
          .get(`http://localhost:5000/api/notification/${user.id}`)
          .then((response) => {
            io.to(user.socketId || "").emit(
              "getAllNotifications",
              response.data
            );
          })
          .catch((error) => {
            console.error("Error", error);
          });
      }
    });

    socket.on(
      "sendNotification",
      ({ senderUser, receiverAuthor, type }: INotificationObject) => {
        const receiver = getUser(receiverAuthor.id || "");
        if (receiver && receiver?.socketId) {
          io.to(receiver?.socketId).emit("updateNotifications", {
            senderUser,
            receiverAuthor,
            type,
          });
          axios
            .post(
              "http://localhost:5000/api/notification",
              {
                senderUser,
                receiverAuthor,
                type,
              },
              {
                headers: {
                  "Content-Type": "application/json",
                },
              }
            )
            .then((response) => {
              io.to(receiver.socketId || "").emit(
                "updateNotification",
                response.data.notifications
              );
            })
            .catch((error) => {
              console.error("Error", error);
            });
        }
      }
    );

    socket.on("disconnect", () => {
      removeUser(socket.id);
    });
  });

  return io;
}
