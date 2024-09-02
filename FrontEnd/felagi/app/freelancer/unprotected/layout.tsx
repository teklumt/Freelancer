import { SignedOut, SignedIn } from "@clerk/nextjs";
import Navbar from "./components/Navbar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="p-4">
        <header className="w-full">
          <Navbar></Navbar>
        </header>
        <main>{children}</main>
      </body>
    </html>
  );
}
