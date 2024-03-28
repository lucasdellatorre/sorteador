'use client';

import { getEdicao, novaEdicao } from "@/app/lib/aposta";
import { Button } from "../button";
import { lusitana } from "../fonts";
import { useState } from "react";

export default function NovaEdicao({ aEdicao }: any) {
  const [edicao, setEdicao] = useState(aEdicao)

  const handleNovaEdicao = async () => {
    try {
      setEdicao(edicao+1)
      const raffle = await novaEdicao();
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div className="flex flex-col justify-center items-center min-h-screen">
      <h1 className={`${lusitana.className} text-4xl md:text-6xl mb-4`}>
        Edição {edicao}

      </h1>
      <div className="flex flex-col items-center gap-4">
        <h2 className="mb-2 text-lg">Deseja iniciar uma nova edição?</h2>
        <Button
          className="text-lg px-8 py-4"
          onClick={handleNovaEdicao}
          type="submit"
        >
          Nova edição
        </Button>
      </div>
    </div>
  );
}