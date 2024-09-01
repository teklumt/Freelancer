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
  MessageCircleQuestion,
  MessageSquareText,
  NotebookTabs,
  Search,
  Settings,
  UserRoundPen,
  Users,
} from "lucide-react";
import Link from "next/link";

const Sidebar = () => {
  const menuList = [
    {
      group: "GENERAL",
      items: [
        {
          link: "/freelancer/protected/dashboard",
          icon: <LayoutDashboard />,
          text: "Dashboard",
        },
        {
          link: "/freelancer/protected/messages",
          icon: <MessageSquareText />,
          text: "Messages",
        },
        {
          link: "/freelancer/protected/my-applications",
          icon: <NotebookTabs />,
          text: "My Applications",
        },
        {
          link: "/freelancer/protected/find-jobs",
          icon: <Search />,
          text: "Find Jobs",
        },
        {
          link: "/freelancer/protected/browse-companies",
          icon: <Building2 />,
          text: "Browse Companies",
        },
        {
          link: "/freelancer/protected/public-profile",
          icon: <UserRoundPen />,
          text: "My Public Profile",
        },
      ],
    },
    {
      group: "SETTINGS",
      items: [
        {
          link: "/freelancer/protected/settings",
          icon: <Settings />,
          text: "Settings",
        },
        {
          link: "/freelancer/protected/help-center",
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
      <div>
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
    </div>
  );
};

export default Sidebar;
