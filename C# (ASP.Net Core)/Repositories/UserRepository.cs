using System.Collections.Generic;
using System.Data;
using System.Threading.Tasks;
using Dapper;
using MySql.Data.MySqlClient;
using WebCRUD.Models;
using Dumpify;

namespace WebCRUD.Repositories;

public class UserRepository
{
	private readonly IDbConnection _db;

	public UserRepository(IDbConnection db)
	{
		_db = db;
	}

	public async Task<IEnumerable<UserModel>> GetAll()
	{
		var sql = @"SELECT *,
						   created_at AS CreatedAt,
						   updated_at AS UpdatedAt
					FROM users";
		return await _db.QueryAsync<UserModel>(sql);
	}

	public async Task<UserModel> GetByID(int Id )
	{
		var sql = @"SELECT *,
						   created_at AS CreatedAt,
						   updated_at AS UpdatedAt
					FROM users
					WHERE id=@Id";
		return await _db.QuerySingleOrDefaultAsync<UserModel>(sql, new {Id=Id});
	}

	public async Task<int> Insert(UserModel user)
	{
		var sql = @"INSERT INTO users (name, email, website)
					VALUES (@Name, @Email, @Website);
					SELECT LAST_INSERT_ID();";

		return await _db.ExecuteScalarAsync<int>(sql, user);
	}

	public async Task<int> Update(UserModel user)
	{
		var sql = @"
				UPDATE users
				SET name = @Name,
					email = @Email,
					website = @Website
				WHERE id=@Id";
		return await _db.ExecuteAsync(sql, user);
	}

	public async Task<int> Delete(int Id)
	{
		var sql = @"
				DELETE FROM users
				WHERE id=@Id";
		return await _db.ExecuteAsync(sql, new {Id=Id});
	}
}