import { Aposta, CreateAposta, Raffle, RaffleStats } from "./definitions";

export async function getApostas() {
  const res = await fetch(`http://localhost:9999/gamble/list`, { cache: 'no-store' });
  const data: Aposta = await res.json();

  return data;
}

export async function getEdicao() {
  const res = await fetch(`http://localhost:9999/raffle/last`, { cache: 'no-store' });
  const data: Raffle = await res.json();

  return data;
}

export async function novaEdicao() {
  const res = await fetch(`http://localhost:9999/raffle/create`, { cache: 'no-store', method: 'POST' });
  const data: Raffle = await res.json();

  return data;
}


export async function iniciarSorteio() {
  const res = await fetch(`http://localhost:9999/raffle/start`, { 
    cache: 'no-store', 
    method: 'POST', 
  });
  return res
}

export async function gerarSorteio() {
  const res = await fetch(`http://localhost:9999/raffle/generate`, { 
    cache: 'no-store', 
    method: 'POST', 
  });

  const data: Raffle = await res.json()

  return data
}

export async function fecharSorteio() {
  const res = await fetch(`http://localhost:9999/raffle/clsoe`, { 
    cache: 'no-store', 
    method: 'POST', 
  });
  return res
}

export async function createAposta({name, cpf, numbers }: CreateAposta) {
  const res = await fetch(`http://localhost:9999/gamble/create`, { 
    cache: 'no-store', 
    method: 'POST', 
    body: JSON.stringify({
      name,
      cpf,
      numbers
    }), 
    headers: {
    'Content-Type': 'application/json',
  }});
  return res
}

export async function getApuracao() {
  const res = await fetch(`http://localhost:9999/raffle/stats`, { 
    cache: 'no-store', 
    method: 'GET', 
  });

  const data: RaffleStats = await res.json()

  return data
}

export async function getGambleCount() {
  const res = await fetch(`http://localhost:9999/raffle/count`, { 
    cache: 'no-store', 
    method: 'GET', 
  });

  const data: number[] = await res.json()

  console.log(data)
  return data
}