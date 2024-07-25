import { Module } from '@nestjs/common';
import { EventsModule } from '../../partner1/src/events/events.module';
import { SpotsModule } from '../../partner1/src/spots/spots.module';
import { ConfigModule } from '@nestjs/config';
import { PrismaModule } from '@app/core/prisma/prisma-core.module';
import { AuthModule } from '@app/core/auth/auth.module';

@Module({
  imports: [
    ConfigModule.forRoot({envFilePath: '.env.partner1', isGlobal: true}),
    AuthModule,
    PrismaModule, 
    EventsModule, 
    SpotsModule,
  ],
})
export class Partner1Module {}
