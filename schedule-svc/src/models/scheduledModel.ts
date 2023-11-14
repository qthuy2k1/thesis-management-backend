import * as admin from "firebase-admin";
import { ICouncilDef, IThesisDef } from "../interface/scheduleDef";

export class ScheduleDefModel {
  private static db = admin.firestore();
  private static scheduleDoc = this.db.collection("schedule-defs");

  // SAVE SCHEDULE TO DB
  static async saveScheduleDef(thesis: IThesisDef): Promise<IThesisDef> {
    // Delete all documents in the table
    await this.deleteAllDocuments();

    // Add a new document
    const docRef = await this.scheduleDoc.add({
      ...thesis,
      createdAt: admin.firestore.Timestamp.now(),
    });

    const createdScheduleSnapshot = await docRef.get();
    const createdSchedule = {
      id: createdScheduleSnapshot.id,
      ...createdScheduleSnapshot.data(),
    } as IThesisDef;

    return createdSchedule;
  }

  static async deleteAllDocuments(): Promise<void> {
    const querySnapshot = await this.scheduleDoc.get();

    const batch = admin.firestore().batch();
    querySnapshot.forEach((doc) => {
      batch.delete(doc.ref);
    });

    await batch.commit();
  }

  // GET SCHEDULE
  static async getSchedule(): Promise<IThesisDef> {
    const docRef = await this.scheduleDoc.get();
    const snapshot = docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as IThesisDef)
    );
    return snapshot[0];
  }
  // GET ONE SCHEDULE
  static async getCouncilInSchedule(id: string): Promise<ICouncilDef | null> {
    const docRef = await this.getSchedule();
    const thesis = docRef.thesis.find((item: any) => item.id === id);
    if (!thesis) {
      return null;
    }
    return { ...thesis };
  }
  // GET SCHEDULE FOR STUDENT
  static async getScheduleForStudent(id: string): Promise<ICouncilDef | null> {
    const docRef = await this.getSchedule();
    const thesis = docRef.thesis.find((item: ICouncilDef) =>
      item.schedule.timeSlots.find(
        (timeSlot) => timeSlot.student.infor.id === id
      )
    );
    const studentSchedule = thesis?.schedule.timeSlots.filter(
      (item) => item.student.infor.id === id
    );

    if (!thesis) {
      return null;
    }
    return {
      ...thesis,
      schedule: {
        room: thesis.schedule.room,
        timeSlots: studentSchedule || [],
      },
    };
  }

  static async getScheduleForLecturer(
    id: string
  ): Promise<ICouncilDef[] | null> {
    const docRef = await this.getSchedule();
    const thesis = docRef.thesis.filter((item: ICouncilDef) =>
      item.council.find((item) => item.id === id)
    );
    if (!thesis) {
      return null;
    }
    return thesis;
  }
}
