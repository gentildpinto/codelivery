import { Document } from 'mongoose';
import { Prop, Schema, raw, SchemaFactory } from '@nestjs/mongoose';

export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  title: string;

  @Prop(
    raw({
      latitude: { type: Number },
      longitude: { type: Number },
    }),
  )
  startPosition: {
    latitude: number;
    longitude: number;
  };

  @Prop(
    raw({
      latitude: { type: Number },
      longitude: { type: Number },
    }),
  )
  endPosition: {
    latitude: number;
    longitude: number;
  };
}

export const RouteSchema = SchemaFactory.createForClass(Route);
