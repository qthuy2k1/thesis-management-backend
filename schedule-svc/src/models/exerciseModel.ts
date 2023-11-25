import * as admin from "firebase-admin";
import { IExerciseObject } from "../interface/exercise";

export class ExerciseModel {
  private static db = admin.firestore();
  private static exerciseDoc = this.db.collection("exercises");
  // CREATE EXERCISE
  static async createExercise(
    exercise: Omit<IExerciseObject, "id">
  ): Promise<IExerciseObject> {
    const existingExercise = await this.exerciseDoc
      .where("uid", "==", exercise.uid)
      .get();
    if (!existingExercise.empty) {
      throw new Error("Exercise exists");
    } else {
      const docRef = await this.exerciseDoc.add({
        ...exercise,
        createdAt: admin.firestore.Timestamp.now(),
      });
      const createdStageSnapshot = await docRef.get();
      const createdStage = {
        id: createdStageSnapshot.id,
        ...createdStageSnapshot.data(),
      } as IExerciseObject;
      return createdStage;
    }
  }

  // GET ONE EXERCISE
  static async getExercise(id: string): Promise<IExerciseObject | null> {
    const docRef = await this.exerciseDoc.doc(id).get();
    if (!docRef.exists) {
      return null;
    }
    const data = docRef.data() as IExerciseObject;
    return { ...data };
  }

  // GET ALL EXERCISE
  static async getAllExercise(): Promise<IExerciseObject[]> {
    const docRef = await this.exerciseDoc.orderBy("createdAt", "desc").get();
    return docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as IExerciseObject)
    );
  }

  // GET ALL EXERCISE IN CLASS
  static async getAllExerciseInClass(id: string): Promise<IExerciseObject[]> {
    const docRef = await this.exerciseDoc.where("classroom.id", "==", id).get();
    return docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as IExerciseObject)
    );
  }

  // GET ALL EXERCISE IN REPORT STAGE
  static async getAllExerciseInReportStage(
    classroomId: string,
    categoryId: string
  ): Promise<IExerciseObject[]> {
    const docRef = await this.exerciseDoc
      .where("classroom.id", "==", classroomId)
      .where("category.id", "==", categoryId)
      .get();
    return docRef.docs.map(
      (doc) => ({ id: doc.id, ...doc.data() } as IExerciseObject)
    );
  }

  // UPDATE EXERCISE
  static async updateExercise(exercise: IExerciseObject): Promise<void> {
    const { id, ...docRef } = exercise;
    await this.exerciseDoc.doc(id).update(docRef);
  }
  // DELETE EXERCISE
  static async deleteExercise(id: string): Promise<IExerciseObject[]> {
    await this.exerciseDoc.doc(id).delete();
    const snapshot = await this.exerciseDoc.where("id", "==", id).get();
    const exercises: IExerciseObject[] = [];
    snapshot.forEach((doc) => {
      exercises.push({ id: doc.id, ...doc.data() } as IExerciseObject);
    });
    return exercises;
  }
}