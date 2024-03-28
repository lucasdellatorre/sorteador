'use client'

import { ArrowPathIcon } from '@heroicons/react/24/outline';
import clsx from 'clsx';
import { lusitana } from '@/app/ui/fonts';
import { Aposta } from '@/app/lib/definitions';
import { getApostas, getEdicao } from '@/app/lib/aposta';
import { useEffect, useState } from 'react';



export default function LatestGambles() {
  const [apostas, setApostas] = useState<Aposta[]>([])
  const [loading, setLoading] = useState<boolean>(false)

  const setPageState = async () => {
    try {
      setLoading(true)
      const data: Aposta[] = await getApostas()
      if (data) {
        setApostas(data)
        console.log(data)
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }

  }

  useEffect(() => {
    setPageState()
  }, []);

  return (
    <div className="flex w-full flex-col md:col-span-4">
      <h2 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Todas as apostas
      </h2>
      <div className="flex grow flex-col justify-between rounded-xl bg-gray-50 p-4">
        <div className="bg-white px-6">
          {
            apostas && apostas.length !== 0 ? apostas.map((aposta, i) => {
              return (
                <div
                  className={clsx(
                    'flex flex-row items-center justify-between py-4'
                  )}
                >
                  <div className="flex items-center">
                    <div className="min-w-0">
                      <p className="truncate text-sm font-semibold md:text-base">
                        {aposta.name}
                      </p>
                      <p className="hidden text-sm text-gray-500 sm:block">
                        {aposta.cpf}
                      </p>
                      <p className="hidden text-sm text-gray-500 sm:block">
                        Edição {aposta.raffle_id.toString()}
                      </p>
                      <p className="hidden text-sm text-gray-500 sm:block">
                        {aposta.numbers.toString()}
                      </p>
                    </div>
                  </div>
                  <p
                    className={`${lusitana.className} truncate text-sm font-medium md:text-base`}
                  >
                    {aposta.gamble_id}
                  </p>
                </div>
              );
            }) :
              !loading && apostas.length === 0 ? <h1>Não há apostas!</h1> : null
          }
        </div>
        <div className="flex items-center pb-2 pt-6">
          <ArrowPathIcon className="h-5 w-5 text-gray-500" />
          <h3 className="ml-2 text-sm text-gray-500 "> {loading ? "Loading" : "Updated just now"}</h3>
        </div>
      </div>
    </div>
  );
}
