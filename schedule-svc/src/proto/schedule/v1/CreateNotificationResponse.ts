// Original file: src/proto/schedule.proto

import type { Notification as _schedule_v1_Notification, Notification__Output as _schedule_v1_Notification__Output } from '../../schedule/v1/Notification';

export interface CreateNotificationResponse {
  'notification'?: (_schedule_v1_Notification | null);
  'message'?: (string);
  'notifications'?: (_schedule_v1_Notification)[];
}

export interface CreateNotificationResponse__Output {
  'notification'?: (_schedule_v1_Notification__Output);
  'message'?: (string);
  'notifications'?: (_schedule_v1_Notification__Output)[];
}
