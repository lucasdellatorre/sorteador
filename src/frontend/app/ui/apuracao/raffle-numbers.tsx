import clsx from 'clsx';
import { getApuracao } from '@/app/lib/aposta';

export default async function RaffleNumbers() {
    const { numbers } = await getApuracao()

    return (
        <div className="flex row-wrap gap-2">
            {numbers.map((number) => {

                const circleClasses = clsx(
                    'w-12 h-12 rounded-full border border-gray-500 flex justify-center items-center cursor-pointer',
                );

                return (
                    <div
                        key={number}
                        className={circleClasses}
                    >
                        {number}
                    </div>
                );
            })}
        </div>
    );
}
