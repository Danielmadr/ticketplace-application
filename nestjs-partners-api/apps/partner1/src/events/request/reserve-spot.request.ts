import { TicketKind } from '@prisma/client';

export class ReserveSpotRequest {
  spots: string[]; // [A1,A2,A3]
  ticket_kind: TicketKind;
  email: string;
}
