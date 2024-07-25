import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { Observable } from 'rxjs';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private configService: ConfigService) { }

  canActivate(
    context: ExecutionContext,
  ): boolean {
    const request = context.switchToHttp().getRequest<Request>();
    const token = request.headers['x-api-token'];
    console.log(token) 
    return token === this.configService.get('API_TOKEN')
  }
}
