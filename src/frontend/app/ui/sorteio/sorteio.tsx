'use client';

import clsx from 'clsx';
import Link from 'next/link';
import { Button } from '@/app/ui/button';
import { FormEvent, useEffect, useState } from 'react';
import { lusitana } from '../fonts';
import { iniciarSorteio, gerarSorteio, getEdicao } from '@/app/lib/aposta';

export default function Sorteio() {
    const [rodada, setRodada] = useState<number>(1)
    const [isRaffleOpenForGamble, setIsRaffleOpenForGamble] = useState<boolean>(true);
    const [selectedNumbers, setSelectedNumbers] = useState<number[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [findingWinners, setFindingWinners] = useState<boolean>(false);
    const [isRaffleFinished, setIsRaffleFinished] = useState<boolean>(false);

    const setPageState = async () => {
        setLoading(true)
        const { raffle, active } = await getEdicao()
        setLoading(false)
        if (!active) {
            setIsRaffleOpenForGamble(false)
            setSelectedNumbers(raffle)
            setRodada(raffle.length - 4)
            setIsRaffleFinished(true)
        }
    }

    useEffect(() => {
        setPageState()
    }, []);

    const openRaffle = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            await iniciarSorteio() // isso cancela as apostas
        } catch (error: any) {
            console.log(error)
        } finally {
            setIsRaffleOpenForGamble(false)
        }
    }

    const sortear = async () => {
        setSelectedNumbers([]);
        setFindingWinners(true)
        const sorteio = await gerarSorteio()

        if (sorteio === null) {
            return
        }

        for (let i = 0; i < 5; i++) {
            setSelectedNumbers(prevSelectedNumbers => [...prevSelectedNumbers, sorteio.raffle[i]]);
            await new Promise(resolve => setTimeout(resolve, 600));
        }

        for (let i = 5; i < sorteio.raffle.length; i++) {
            setRodada(i - 4)
            setSelectedNumbers(prevSelectedNumbers => [...prevSelectedNumbers, sorteio.raffle[i]]);
            await new Promise(resolve => setTimeout(resolve, 1000));
        }
        setFindingWinners(false)
        setIsRaffleFinished(true)
    };

    return (
        <div>
            {loading ? (
                <div> loading </div>
            ) : (
                <>
                    {!isRaffleOpenForGamble ? (
                        <div className="flex flex-col justify-center items-center min-h-screen">
                            <h1 className={`${lusitana.className} text-4xl md:text-6xl gap-4`}> Sorteio </h1>
                            <h2 className={`${lusitana.className} text-4xl md:text-2xl mt-6`}> Rodada {rodada} </h2>
                            <div className="flex justify-content center">
                                <div className="grid grid-cols-10 gap-8 mt-20">
                                    {[...Array(50)].map((_, index) => {
                                        const number = index + 1;
                                        const isSelected = selectedNumbers.includes(number);

                                        const circleClasses = clsx(
                                            'w-16 h-16 rounded-full border border-gray-500 flex justify-center items-center cursor-pointer',
                                            {
                                                'bg-red-500 text-white': isSelected,
                                            }
                                        );

                                        return (
                                            <div
                                                key={number}
                                                className={circleClasses}
                                                style={{ animationDelay: `${index * 0.1}s` }}

                                            >
                                                {number}
                                            </div>
                                        );
                                    })}
                                </div>
                            </div>
                            {isRaffleFinished ? (
                                <Link
                                    href="/home/apuracao"
                                  className="text-lg px-8 py-4 mt-10  rounded-lg bg-red-500 text-sm font-medium text-white transition-colors hover:bg-red-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-500 active:bg-red-600 aria-disabled:cursor-not-allowed aria-disabled:opacity-50"
                                >
                                    Apuração
                                </Link>

                            ): (
                                <Button
                                className="text-lg px-16 py-6 mt-10"
                                onClick={() => sortear()}
                                disabled={findingWinners}
                            >
                                {findingWinners ? "Procurando Vencedores..." : "Sortear"}
                            </Button>
                            )}
                        </div>
                    ) :
                        (
                            <div className="flex flex-col justify-center items-center min-h-screen">
                                <h1 className={`${lusitana.className} text-4xl md:text-6xl mb-4`}>
                                    Iniciar Sorteio
                                </h1>
                                <div className="flex flex-col items-center gap-4">
                                    <h2 className="mb-2 text-lg">Deseja iniciar o sorteio? (apostas não serão mais permitidas) </h2>
                                    <Button
                                        className="text-lg px-8 py-4"
                                        onClick={openRaffle}
                                    >
                                        Iniciar sorteio
                                    </Button>
                                </div>
                            </div>
                        )
                    }
                </>
            )}
        </div>
    );
}
