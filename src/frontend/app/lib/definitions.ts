// This file contains type definitions for your data.
// It describes the shape of the data, and what data type each property should accept.
// For simplicity of teaching, we're manually defining these types.
// However, these types are generated automatically if you're using an ORM such as Prisma.

export type RaffleStats = {
  numbers: number[];
  rounds: number;
  winners_count: number;
  winners: Winner[];
}

export type GambleStats = {
  number: number;
  count: number;
}

export type Winner = {
  gamble_id: number
  name: string
  cpf: string
  numbers: number[]
  raffle_id: number
  prize: number
}

export type Raffle = {
  raffle_id: number;
  raffle: number[];
  active: boolean;
  prize: number;
}

export type RaffleNumber = {
  value: string;
  isSelected: boolean;
}

export type Aposta = {
  gamble_id: number
  name: string
  cpf: string
  numbers: number[]
  raffle_id: number
}

export type CreateAposta = {
  name: string
  cpf: string
  numbers: number[]
}