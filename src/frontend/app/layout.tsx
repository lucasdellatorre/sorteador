import '@/app/ui/global.css';
import { inter } from '@/app/ui/fonts';
// import {NextUIProvider} from "@nextui-org/system";
 
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${inter.className} antialiased`}>
        {/* <NextUIProvider> */}
        {children}
        {/* </NextUIProvider> */}
      </body>
    </html>
  );
}