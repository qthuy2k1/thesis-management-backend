// Original file: src/proto/schedule.proto

import type { AuthorExerciseResponse as _schedule_v1_AuthorExerciseResponse, AuthorExerciseResponse__Output as _schedule_v1_AuthorExerciseResponse__Output } from '../../schedule/v1/AuthorExerciseResponse';
import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../../google/protobuf/Timestamp';
import type { Long } from '@grpc/proto-loader';

export interface ClassroomResponse {
  'id'?: (number | string | Long);
  'status'?: (string);
  'lecturer'?: (_schedule_v1_AuthorExerciseResponse | null);
  'classCourse'?: (string);
  'topicTags'?: (string);
  'quantityStudent'?: (number | string | Long);
  'createdAt'?: (_google_protobuf_Timestamp | null);
  'updatedAt'?: (_google_protobuf_Timestamp | null);
  '_topicTags'?: "topicTags";
}

export interface ClassroomResponse__Output {
  'id'?: (Long);
  'status'?: (string);
  'lecturer'?: (_schedule_v1_AuthorExerciseResponse__Output);
  'classCourse'?: (string);
  'topicTags'?: (string);
  'quantityStudent'?: (Long);
  'createdAt'?: (_google_protobuf_Timestamp__Output);
  'updatedAt'?: (_google_protobuf_Timestamp__Output);
}
