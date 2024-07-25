import { Controller, Get, Post, Body, Patch, Param, Delete, HttpCode, UseGuards } from '@nestjs/common';
import { EventsService } from '@app/core/events/events.service';
import { CreateEventReques} from './request/create-event.request';
import { UpdateEventRequest } from './request/update-event.request';
import { ReserveSpotRequest } from './request/reserve-spot.request';
import { AuthGuard } from '@app/core/auth/auth.guard';


@Controller('events')
export class EventsController {
  constructor(private readonly eventsService: EventsService) {}

// NOTE: DTO stands for Data Transfer Object, a dummy object used only to carry data.
  @UseGuards(AuthGuard)
  @HttpCode(201)
  @Post()
  create(@Body() createEventRequest: CreateEventReques) {
    return this.eventsService.create(createEventRequest);
  }

  @UseGuards(AuthGuard)
  @HttpCode(201)
  @Post(':id/reserve')
  createReservation(@Body() dto: ReserveSpotRequest, @Param('id') eventId: string) {
    return this.eventsService.reserveSpot({...dto, eventId});
  }


  @UseGuards(AuthGuard)
  @HttpCode(200)
  @Get()
  findAll() {
    return this.eventsService.findAll();
  }


  @UseGuards(AuthGuard)
  @HttpCode(200)
  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.eventsService.findOne(id);
  }


  @UseGuards(AuthGuard)
  @HttpCode(200)
  @Patch(':id')
  update(@Param('id') id: string, @Body() updateEventRequest: UpdateEventRequest) {
    return this.eventsService.update(id, updateEventRequest);
  }


  @UseGuards(AuthGuard)
  @HttpCode(204)
  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.eventsService.remove(id);
  }

}


