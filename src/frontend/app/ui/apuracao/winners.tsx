import clsx from 'clsx';
import { lusitana } from '@/app/ui/fonts';
import { getApuracao } from '@/app/lib/aposta';
export default async function WinnersList() {
    let {winners} = await getApuracao()

    if (winners != null && winners.length != 0) {
        winners = winners.sort((a, b) => a.name.localeCompare(b.name))
    } else {
        winners = []
    }

  return (
    <div className="flex w-full flex-col md:col-span-2">
      <h2 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Apostas Vencedoras
      </h2>
      <div className="flex grow flex-col justify-between rounded-xl bg-gray-50 p-4">
        <div className="bg-white px-6">
          {winners !== null && winners.length !== 0 ? winners.map((invoice, i) => {
            return (
              <div
                className={clsx(
                  'flex flex-row items-center justify-between py-4'
                )}
              >
                <div className="flex items-center">
                  <div className="min-w-0">
                    <p className="truncate text-sm font-semibold md:text-base">
                      {invoice.name}
                    </p>
                    <p className="hidden text-sm text-gray-500 sm:block">
                      {invoice.cpf}
                    </p>
                    <p className="hidden text-sm text-gray-500 sm:block">
                      Edição {invoice.raffle_id.toString()}
                    </p>
                    <p className="hidden text-sm text-gray-500 sm:block">
                      {invoice.numbers.toString()}
                    </p>
                    <p className="hidden text-sm text-gray-500 sm:block">
                       Prêmio: R${invoice.prize.toString()}
                    </p>
                  </div>
                </div>
                <p
                  className={`${lusitana.className} truncate text-sm font-medium md:text-base`}
                >
                  {invoice.gamble_id}
                </p>
              </div>
            );
          }) :  (
                  <p className="hidden text-sm text-gray-500 sm:block">
                    Não houveram vencedores!
                </p>
          ) 
        }
        </div>
      </div>
    </div>
  );
}
