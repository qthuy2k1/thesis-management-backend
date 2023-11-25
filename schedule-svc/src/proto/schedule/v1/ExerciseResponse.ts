// Original file: src/proto/schedule.proto

import type { ClassroomResponse as _schedule_v1_ClassroomResponse, ClassroomResponse__Output as _schedule_v1_ClassroomResponse__Output } from '../../schedule/v1/ClassroomResponse';
import type { ReportingStageExerciseResponse as _schedule_v1_ReportingStageExerciseResponse, ReportingStageExerciseResponse__Output as _schedule_v1_ReportingStageExerciseResponse__Output } from '../../schedule/v1/ReportingStageExerciseResponse';
import type { AuthorExerciseResponse as _schedule_v1_AuthorExerciseResponse, AuthorExerciseResponse__Output as _schedule_v1_AuthorExerciseResponse__Output } from '../../schedule/v1/AuthorExerciseResponse';
import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../../google/protobuf/Timestamp';
import type { AttachmentExerciseResponse as _schedule_v1_AttachmentExerciseResponse, AttachmentExerciseResponse__Output as _schedule_v1_AttachmentExerciseResponse__Output } from '../../schedule/v1/AttachmentExerciseResponse';

export interface ExerciseResponse {
  'id'?: (string);
  'title'?: (string);
  'description'?: (string);
  'classroomID'?: (_schedule_v1_ClassroomResponse | null);
  'deadline'?: (string);
  'category'?: (_schedule_v1_ReportingStageExerciseResponse | null);
  'author'?: (_schedule_v1_AuthorExerciseResponse | null);
  'createdAt'?: (_google_protobuf_Timestamp | null);
  'updatedAt'?: (_google_protobuf_Timestamp | null);
  'attachments'?: (_schedule_v1_AttachmentExerciseResponse)[];
}

export interface ExerciseResponse__Output {
  'id'?: (string);
  'title'?: (string);
  'description'?: (string);
  'classroomID'?: (_schedule_v1_ClassroomResponse__Output);
  'deadline'?: (string);
  'category'?: (_schedule_v1_ReportingStageExerciseResponse__Output);
  'author'?: (_schedule_v1_AuthorExerciseResponse__Output);
  'createdAt'?: (_google_protobuf_Timestamp__Output);
  'updatedAt'?: (_google_protobuf_Timestamp__Output);
  'attachments'?: (_schedule_v1_AttachmentExerciseResponse__Output)[];
}
