import { getApuracao, getGambleCount } from '@/app/lib/aposta';
import WinnersList from './winners';
import GamblersTable from '../gamblers/table';
import RaffleNumbers from './raffle-numbers';
import { lusitana } from '../fonts';
import { RaffleStats } from '@/app/lib/definitions';
import Link from 'next/link';
import { Card } from '../dashboard/cards';

export default async function Wraper() {
    function parse(arr: number[]) {
        const data = []
        for (let i = 0; i < arr.length; i++) {
            if (arr[i] != 0) {
                data.push({ number: i, count: arr[i] })
            }
        }
        data.sort((a, b) => b.count - a.count)
        console.log(data)
        return data
    }

    const raffle: RaffleStats = await getApuracao()
    const data = await getGambleCount()
    return (
        <>
            {raffle.numbers === null ? (
                <div className="flex flex-col justify-center items-center min-h-screen">
                    <h1 className={`${lusitana.className} text-4xl md:text-6xl mb-4`}>
                        Apuração indisponível
                    </h1>
                    <div className="flex flex-col items-center gap-4">
                        <h2 className="mb-2 text-lg">Deseja iniciar um novo sorteio?</h2>
                        <Link
                            className="text-lg px-8 py-4 flex h-10 items-center rounded-lg bg-red-500 text-sm font-medium text-white transition-colors hover:bg-red-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-500 active:bg-red-600 aria-disabled:cursor-not-allowed aria-disabled:opacity-50"
                            href="/home/sorteio"
                            type="submit"
                        >
                            Novo Sorteio
                        </Link>
                    </div>
                </div>
            ) : (
                <div>
                    <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-4">
                        <Card title="Rodadas" value={raffle.rounds} type="pending" />
                        <Card title="Apostas Vencedoras" value={raffle.winners_count} type="collected" />
                    </div>
                    <h2 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
                        Números sorteados
                    </h2>
                    <RaffleNumbers />
                    <div className="mt-6 grid grid-cols-1 gap-6 md:grid-cols-4 md:grid-cols-3">
                        <WinnersList />
                        <GamblersTable gambleNumbers={parse(data)} />
                    </div>
                </div>
            )}

        </>


    );
}
