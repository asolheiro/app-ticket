import { Controller, Get, Post, Body, Patch, Param, Delete, UseGuards } from '@nestjs/common';
import { SpotsService } from '@app/core/spots/spots.service';
import { CreateSpotRequest } from './request/create-spot.request';
import { UpdateSpotRequest} from './request/update-spot.request';
import { AuthGuard } from '@app/core/auth/auth.guard';


@Controller('events/:eventId/spots')
export class SpotsController {
  constructor(private readonly spotsService: SpotsService) {}

  @UseGuards(AuthGuard) 
  @Post()
  create(
    @Body() createSpotRequest: CreateSpotRequest, 
    @Param('eventId') eventId: string,
  ) {
    console.log(eventId)
    return this.spotsService.create({
      ...createSpotRequest,
      eventId,
    });
  }

  @UseGuards(AuthGuard)
  @Get()
  findAll(@Param('eventId') eventId: string) {
    return this.spotsService.findAll(eventId);
  }

  @Get(':spotId')
  @UseGuards(AuthGuard)
  findOne(@Param('spotId') spotId: string, @Param('eventId') eventId: string) {
    return this.spotsService.findOne(spotId, eventId);
  }

  @UseGuards(AuthGuard)
  @Patch(':spotId')
  update(@Param('spotId') spotId: string, @Param('eventId') eventId: string, @Body() updateSpotRequest: UpdateSpotRequest) {
    return this.spotsService.update(spotId, eventId, updateSpotRequest,);
  }

  @UseGuards(AuthGuard)
  @Delete(':spotId')
  remove(@Param('spotId') spotId: string, @Param('eventId') eventId: string) {
    return this.spotsService.remove(spotId, eventId);
  }
}
