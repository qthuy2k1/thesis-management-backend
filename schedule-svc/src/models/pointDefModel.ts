import * as admin from "firebase-admin";
import { IAssessItem, IPointDefObject } from "../interface/pointDef";

export class PointDefModel {
  private static db = admin.firestore();
  private static pointDoc = this.db.collection("point-def");

  // CREATE POINT DEFENSE
  static async createPointDef(
    point: Omit<IPointDefObject, "id">
  ): Promise<IPointDefObject> {
    const existingUser = await this.pointDoc
      .where("student.id", "==", point.student.id)
      .get();
    if (!existingUser.empty) {
      throw new Error("PointDef exists");
    } else {
      const docRef = await this.pointDoc.add({
        ...point,
        createdAt: admin.firestore.Timestamp.now(),
      });
      const createdStageSnapshot = await docRef.get();
      const createdStage = {
        id: createdStageSnapshot.id,
        ...createdStageSnapshot.data(),
      } as IPointDefObject;
      return createdStage;
    }
  }

  // UPDATE POINT DEFENSE
  static async updatePointDef(point: IPointDefObject): Promise<void> {
    const { id, assesses, ...docRef } = point;
    const existingPointDef = await this.pointDoc.doc(id).get();

    if (!existingPointDef.exists) {
      throw new Error("PointDef does not exist");
    }
    const updatedAssesses = [...existingPointDef.data()?.assesses, ...assesses];
    await this.pointDoc
      .doc(id)
      .update({ ...docRef, assesses: updatedAssesses });
  }

  static async createOrUpdatePointDef(
    point: Omit<IPointDefObject, "id">
  ): Promise<IPointDefObject> {
    const studentId = point.student.id;
    const existingPointDef = await this.pointDoc
      .where("student.id", "==", studentId)
      .get();

    if (!existingPointDef.empty) {
      const pointDefId = existingPointDef.docs[0].id;
      await this.updatePointDef({ ...point, id: pointDefId });
      return { ...point, id: pointDefId };
    } else {
      return await this.createPointDef(point);
    }
  }

  // GET ONE POINT DEFENSE
  static async getPointDef(id: string): Promise<IPointDefObject | null> {
    const querySnapshot = await this.pointDoc
      .where("student.id", "==", id)
      .get();
    if (querySnapshot.empty) {
      return null;
    }
    const docRef = querySnapshot.docs[0];
    const data = docRef.data() as IPointDefObject;
    return { ...data, id: docRef.id };
  }

  static async getPointDefForLecturer(
    studefId: string,
    lecId: string
  ): Promise<IAssessItem | null> {
    const querySnapshot = await this.pointDoc
      .where("student.id", "==", studefId)
      .get();
    if (querySnapshot.empty) {
      return null;
    }
    const docRef = querySnapshot.docs[0];
    const data = docRef.data() as IPointDefObject;

    const objectWithLecId = data.assesses.find((obj) => obj.lecturer.id === lecId);
    if (objectWithLecId) {
      return { ...objectWithLecId };
    }

    return null;
  }

  // GET ALL POINT DEFENSE
  static async getAllPointDef(): Promise<IPointDefObject[]> {
    const docRef = await this.pointDoc.orderBy("createdAt", "desc").get();
    return docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as IPointDefObject)
    );
  }

  // DELETE POINT DEFENSE
  static async deletePointDef(id: string): Promise<IPointDefObject[]> {
    await this.pointDoc.doc(id).delete();
    const snapshot = await this.pointDoc.where("id", "==", id).get();
    const points: IPointDefObject[] = [];
    snapshot.forEach((doc) => {
      points.push({ id: doc.id, ...doc.data() } as IPointDefObject);
    });
    return points;
  }
}
