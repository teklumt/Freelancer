import Navbar from "../components/Navbar";
import Sidebar from "../components/Sidebar";
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {

  return (
    <>
      <html lang="en">
        <body>
          <div className="flex justify-between items-start">
            <aside>
              <Sidebar />
            </aside>
            <main className="flex flex-col w-full h-full pl-[300px]">
              <header>
                <Navbar />
              </header>
              <div>{children}</div>
            </main>
          </div>
        </body>
      </html>
    </>
  );
}
