import { lusitana } from "@/app/ui/fonts";
import Form from "@/app/ui/registra-aposta/create-form";

export default async function Page() {
  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-4xl`}>
        Registrar Aposta
      </h1>
      <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-4">
      </div>
      <div className="mt-6 grid  gap-6 md:grid-cols-2 lg:grid-cols-16">
        <Form />
        {/* <Image
          src="/tigrinho.jpg"
          width={600}
          height={760}
          className="hidden md:block"
          alt="Tigrinho"
        /> */}
      </div>
    </main>
  );
  }