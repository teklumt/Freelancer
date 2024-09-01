import React from "react";
import UserItem from "./UserItem";
import {
  Command,
  CommandGroup,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Building2,
  CalendarCheck,
  LayoutDashboard,
  LogOut,
  MessageCircleQuestion,
  MessageSquareText,
  NotebookTabs,
  Settings,
  Users,
} from "lucide-react";
import { SignOutButton } from "@clerk/nextjs";

import Link from "next/link";
import { Button } from "@/components/ui/button";

const Sidebar = () => {
  const menuList = [
    {
      group: "GENERAL",
      items: [
        {
          link: "/client/protected/dashboard",
          icon: <LayoutDashboard />,
          text: "Dashboard",
        },
        {
          link: "/client/protected/messages",
          icon: <MessageSquareText />,
          text: "Messages",
        },
        {
          link: "/client/protected/company-profile",
          icon: <Building2 />,
          text: "Company Profile",
        },
        {
          link: "/client/protected/all-applicants",
          icon: <Users />,
          text: "All Applicants",
        },
        {
          link: "/client/protected/job-listing",
          icon: <NotebookTabs />,
          text: "Job Listing",
        },
        {
          link: "/client/protected/my-schedule",
          icon: <CalendarCheck />,
          text: "My Schedule",
        },
      ],
    },
    {
      group: "SETTINGS",
      items: [
        {
          link: "/client/protected/settings",
          icon: <Settings />,
          text: "Settings",
        },
        {
          link: "/client/protected/help-center",
          icon: <MessageCircleQuestion />,
          text: "Help Center",
        },
      ],
    },
  ];

  return (
    <div className="fixed flex flex-col w-[300px] min-w-[300px] h-screen border-r-2 p-4">
      <div className="mb-4">
        <UserItem
          imageSrc="https://github.com/shadcn.png"
          alt="Av"
          header="John Doe"
          text="john.doe@example.com"
        />
      </div>
      <div className="grow">
        <Command>
          <CommandList className="max-h-screen">
            {menuList.map((menu, key) => (
              <CommandGroup key={key} heading={menu.group}>
                {menu.items.map((option, optionKey) => (
                  <CommandItem key={optionKey} asChild>
                    <Link
                      href={option.link}
                      className="flex gap-2 items-center cursor-pointer"
                    >
                      {option.icon}
                      {option.text}
                    </Link>
                  </CommandItem>
                ))}
              </CommandGroup>
            ))}
          </CommandList>
        </Command>
      </div>
      <div>
        <SignOutButton>
          <Button variant={"secondary"}>
            <LogOut className="mr-2 h-4 w-4" /> Logout
          </Button>
        </SignOutButton>
      </div>
    </div>
  );
};

export default Sidebar;
