import { lusitana } from '@/app/ui/fonts';

export default function AcmeLogo() {
  return (
    <div
      className={`${lusitana.className} ml-10 flex flex-col justify-center items-center text-white`}
    >
      
      <p className="text-[44px]">🐯</p>
      <p className="text-[44px]">TigreSena </p>
    </div>
  );
}
