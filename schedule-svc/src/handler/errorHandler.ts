import { Request, Response, NextFunction } from "express";

interface CustomError extends Error {
  message: string;
  statusCode?: number;
}

export const errorHandler = (
  error: CustomError,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const errorMessage = error.message || "Unknown error";
  const statusCode = error.statusCode || 400;

  res.status(400).json({
    status: "Fail to call API",
    message: errorMessage,
  });
};
