'use client';

import clsx from 'clsx';
import { Button } from '@/app/ui/button';
import { FormEvent, useState } from 'react';
import { createAposta } from '@/app/lib/aposta';
import { CreateAposta } from '@/app/lib/definitions';
import { CancelButton } from '../cancel-button';
import { getRandomInt } from '@/app/lib/utils';

export default function Form() {
  const [cpf, setCPF] = useState('');
  const [name, setName] = useState('');
  const [selectedNumbers, setSelectedNumbers] = useState<number[]>([]);
  const [error, setError] = useState<string | null>(null)
  const [finished, setFinished] = useState<boolean>(false)

  function clearFields(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setCPF('')
    setName('')
    setSelectedNumbers([])
    setError(null)
  }

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setError(null) // Clear previous errors when a new request starts
 
    try {
      const aposta: CreateAposta = {cpf, name, numbers: selectedNumbers}

      const res = await createAposta(aposta)

      if (!res.ok) {
        throw new Error('Falha ao registrar aposta. Por favor tente novamente.')
      }

      clearFields(event)
 
    } catch (error: any) {
      setError(error.message)
      console.error(error)
    } 
  }

  const toggleSelection = (number: number) => {
    if (selectedNumbers.includes(number)) {
      setSelectedNumbers(selectedNumbers.filter((n) => n !== number));
    } else {
      if (selectedNumbers.length < 5) {
        setSelectedNumbers([...selectedNumbers, number]);
      } else {
        alert('You can only select up to 5 numbers.');
      }
    }
  };
  const surpresinha = (e: any) => {
    e.preventDefault(); // Prevent default form submission behavior
    setSelectedNumbers([]);
    const used = new Set(selectedNumbers);
  
    for (let count = 0; count < 5; count++) {
      let sortNumber: number;
      do {
        sortNumber = getRandomInt(1, 50);
      } while (used.has(sortNumber));
  
      used.add(sortNumber);
      setSelectedNumbers(prevSelectedNumbers => [...prevSelectedNumbers, sortNumber]);
    }
  };



  return (
    <div>
      <form onSubmit={onSubmit}>
      <div className="rounded-md bg-gray-50 p-20 md:p-6">

        {/* Nome */}
        <div className="mb-4">
          <label htmlFor="amount" className="mb-2 block text-sm font-medium">
            Nome
          </label>
          <div className="relative mt-2 rounded-md">
            <div className="relative">
              <input
                required
                id="nome"
                name="nome"
                type="text"
                step="0.01"
                placeholder="Digite seu nome"
                onChange={(e) => setName(e.target.value)}
                value={name}
                className="peer block w-full rounded-md border border-gray-200 py-2 text-sm outline-2 placeholder:text-gray-500"
              />
            </div>
          </div>
        </div>


        {/* Cpf */}
        <div className="mb-4">
          <label htmlFor="amount" className="mb-2 block text-sm font-medium">
            Cpf
          </label>
          <div className="relative mt-2 rounded-md">
            <div className="relative">
              <input
                required
                id="cpf"
                name="cpf"
                type="text"
                step="0.01"
                placeholder="Digite seu cpf"
                onChange={(e) => setCPF(e.target.value)}
                value={cpf}
                className="peer block w-full rounded-md border border-gray-200 py-2 text-sm outline-2 placeholder:text-gray-500"
              />
            </div>
          </div>
        </div>

        <div className="mb-4">
          <label htmlFor="amount" className="mb-4 block text-sm font-medium">
            Escolha 5 números
          </label>
          <div className="grid grid-cols-10 gap-4">

            {/* Create circles from 1 to 50 */}

            {[...Array(50)].map((_, index) => {
              const number = index + 1;
              const isSelected = selectedNumbers.includes(number);

              const circleClasses = clsx(
                'w-10 h-10 rounded-full border border-gray-500 flex justify-center items-center cursor-pointer',
                {
                  'bg-red-500 text-white': isSelected,
                  'hover:bg-red-500 hover:text-yellow-50': !isSelected // New hover effect class when not selected
                }
              );

              return (
                <div
                  key={number}
                  className={circleClasses}
                  onClick={() => toggleSelection(number)}
                >
                  {number}
                </div>
              );
            })}

            <div className="mt-6  justify-end aligns-items gap-4">
              <Button onClick={(e) => surpresinha(e)}>Surpresinha</Button>
            </div>
          </div>

        </div>

      </div>
      <div className="mt-6 flex gap-4">
        <CancelButton type='reset'
          onClick={clearFields}
          className="flex h-10 items-center rounded-lg bg-gray-100 px-4 text-sm font-medium text-gray-600 transition-colors hover:bg-gray-200"
        >
          Cancelar
        </CancelButton>
        <Button type="submit">Registrar Aposta</Button>
      </div>
    </form>
    {error && <div className='mb-2 block text-sm font-medium text-red-600 mt-4'>{error}</div>}
    {finished && <div className='mb-2 block text-sm font-medium text-green-700 mt-4'> Sucesso ao registrar aposta </div>}
    </div>
  );
}
