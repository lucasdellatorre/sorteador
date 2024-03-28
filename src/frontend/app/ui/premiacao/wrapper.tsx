import { getApuracao, getEdicao } from '@/app/lib/aposta';
import WinnersList from '../apuracao/winners';
import { lusitana } from '../fonts';
import { Raffle, RaffleStats } from '@/app/lib/definitions';
import Link from 'next/link';
import { Card } from '../dashboard/cards';
export default async function Wrpaper() {
    const raffle: RaffleStats = await getApuracao()
    const {prize, raffle_id }: Raffle = await getEdicao()

    return (
        <>
            {raffle.numbers === null ? (
                <div className="flex flex-col justify-center items-center min-h-screen">
                    <h1 className={`${lusitana.className} text-4xl md:text-6xl mb-4`}>
                        Premiação indisponível
                    </h1>
                    <div className="flex flex-col items-center gap-4">
                        <h2 className="mb-2 text-lg">Conclua o sorteio dessa edição</h2>
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
                        <Card title="Edição" value={raffle_id} type="pending" />
                        <Card title="Prêmio acumulado" value={prize} type="collected" />
                    </div>
                    <div className="mt-6 grid grid-cols-1 gap-6 md:grid-cols-4 md:grid-cols-3">
                        <WinnersList />
                    </div>
                </div>
            )}

        </>


    );
}
