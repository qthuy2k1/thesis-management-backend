// Original file: src/proto/schedule.proto

import type { ClassroomResponse as _schedule_v1_ClassroomResponse, ClassroomResponse__Output as _schedule_v1_ClassroomResponse__Output } from '../../schedule/v1/ClassroomResponse';
import type { ReportingStageExerciseResponse as _schedule_v1_ReportingStageExerciseResponse, ReportingStageExerciseResponse__Output as _schedule_v1_ReportingStageExerciseResponse__Output } from '../../schedule/v1/ReportingStageExerciseResponse';
import type { AuthorExerciseResponse as _schedule_v1_AuthorExerciseResponse, AuthorExerciseResponse__Output as _schedule_v1_AuthorExerciseResponse__Output } from '../../schedule/v1/AuthorExerciseResponse';
import type { AttachmentExerciseInput as _schedule_v1_AttachmentExerciseInput, AttachmentExerciseInput__Output as _schedule_v1_AttachmentExerciseInput__Output } from '../../schedule/v1/AttachmentExerciseInput';

export interface ExerciseInput {
  'title'?: (string);
  'description'?: (string);
  'classroom'?: (_schedule_v1_ClassroomResponse | null);
  'deadline'?: (string);
  'categoryID'?: (_schedule_v1_ReportingStageExerciseResponse | null);
  'authorID'?: (_schedule_v1_AuthorExerciseResponse | null);
  'attachments'?: (_schedule_v1_AttachmentExerciseInput)[];
}

export interface ExerciseInput__Output {
  'title'?: (string);
  'description'?: (string);
  'classroom'?: (_schedule_v1_ClassroomResponse__Output);
  'deadline'?: (string);
  'categoryID'?: (_schedule_v1_ReportingStageExerciseResponse__Output);
  'authorID'?: (_schedule_v1_AuthorExerciseResponse__Output);
  'attachments'?: (_schedule_v1_AttachmentExerciseInput__Output)[];
}
