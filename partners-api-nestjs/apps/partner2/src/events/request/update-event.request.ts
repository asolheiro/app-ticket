import { PartialType } from '@nestjs/mapped-types';
import { CreateEventReques } from './create-event.request';

export class UpdateEventRequest extends PartialType(CreateEventReques) {}
