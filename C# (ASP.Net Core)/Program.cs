using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Dumpify;

var builder = WebApplication.CreateBuilder(args);

// Add MVC services
	builder.Services.AddControllersWithViews();
	builder.Services.AddRouting(options =>
	{
		options.LowercaseUrls = true;
		options.AppendTrailingSlash = false; // optional
	});

	var app = builder.Build();

// Middlewares
	if (!app.Environment.IsDevelopment())
	{
		app.UseExceptionHandler("/Home/Error");
		app.UseHsts();
	}

	app.UseHttpsRedirection();
	app.UseStaticFiles();
	app.UseRouting();

	app.UseAuthorization();

// Routes
	app.MapControllerRoute(
		name: "user-list",
		pattern: "users/{action=List}/{id?}",
		defaults: new {controller="User"});

// Default route
	app.MapControllerRoute(
		name: "default",
		pattern: "{controller=User}/{action=Index}/{id?}");


app.Run();