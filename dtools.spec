%define debug_package   %{nil}
%define _build_id_links none
%define _name   dtools
%define _prefix /opt
%define _version 1.000
%define _rel 0
%define _arch x86_64

Name:       dtools
Version:    %{_version}
Release:    %{_rel}
Summary:    dtools

Group:      docker
License:    GPL2.0
URL:        http://git.famillegratton.net:3000/devops/dtools

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: wget,gcc
Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Docker client Swiss knife

%prep
#%setup -q
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_name} .

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
#%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/bin"
#install -Dpm 0755 %{buildroot}/%{_name} "$RPM_BUILD_ROOT%{_prefix}/bin/"
install -Dpm 0755 %{_sourcedir}/%{name} %{buildroot}%{_bindir}/%{name}

%post
strip %{_prefix}/bin/%{name}

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{name}


%changelog

