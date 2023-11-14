import { IRoomDefObject } from "./roomDef";
import { IStudentDef } from "./studentDef";
import { IAuthObject } from "./auth";

export interface IParamSchedule {
  quantityWeek: number;
  startDate: string;
}

export interface IParamAllSchedule {
  startTime: string;
  quantityWeek: number;
  councils: IAuthObject[];
  rooms: IRoomDefObject[];
  studentDefs: IStudentDef[];
}

export interface ITimeSlotItem {
  id: string;
  date: string;
  time: string;
  shift: string;
}

export interface ITimeSlotForStudent {
  timeSlot: ITimeSlotItem;
  student: IStudentDef;
}

export interface IScheduleDefForStudent {
  room: IRoomDefObject;
  timeSlots: ITimeSlotItem[];
}

export interface IScheduleDef {
  room: IRoomDefObject;
  timeSlots: ITimeSlotForStudent[];
}

export interface ICouncilDef {
  id: string;
  council: IAuthObject[];
  schedule: IScheduleDef;
}

export interface IThesisDef {
  id?: string;
  thesis: ICouncilDef[];
  fitness: number;
}
