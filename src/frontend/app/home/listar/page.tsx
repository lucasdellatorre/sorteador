export const dynamic = "force-dynamic"

import LatestGambles from "@/app/ui/dashboard/latest-gambles";

export default async function Page() {
  return (
    <main>
      <div className="mt-6 grid grid-cols-1 gap-6 md:grid-cols-4 lg:grid-cols-8">
        <LatestGambles /> 
      </div>
    </main>
  );
  }