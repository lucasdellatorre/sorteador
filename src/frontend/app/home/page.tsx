import { getEdicao } from "../lib/aposta";
import NovaEdicao from "../ui/novaedicao/nova-edicao";

export default async function Page() {
  const edicao = await getEdicao()
  return (
    <div>
      <NovaEdicao aEdicao={edicao.raffle_id}></NovaEdicao>
    </div>
   
  );
}