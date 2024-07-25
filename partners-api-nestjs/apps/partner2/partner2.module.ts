import { PrismaModule } from '@app/core/prisma/prisma-core.module';
import { Module } from '@nestjs/common'
import { ConfigModule } from '@nestjs/config';
import { SpotsModule } from 'apps/partner2/src/spots/spots.module';
import { EventsModule } from 'apps/partner2/src/events/events.module';
import { AuthGuard } from '@app/core/auth/auth.guard';

@Module({
    imports: [
        ConfigModule.forRoot({ envFilePath: '.env.partner2' }),
        AuthGuard,
        PrismaModule,
        EventsModule,
        SpotsModule,
    ]
}) export class Partner2Module {}